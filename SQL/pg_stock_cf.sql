-- Load data from temporary table
INSERT INTO cf_statement(transfer_day, transfer_time,inflow,outflow, atrribute, check_at )
SELECT CAST(transfer_day AS DATE),
CASE WHEN transfer_time IS NOT NULL THEN CAST(transfer_time AS time without time zone)
ELSE NULL
END as transfer_time,
CASE 
    WHEN inflow IS NULL THEN 0
	ELSE inflow END inflow ,
CASE WHEN outflow IS NULL THEN 0
ELSE outflow END outflow,
CASE 
    WHEN attribute='建行存管' THEN 1
	WHEN attribute='盈利重计' THEN 2
	WHEN attribute='年度计提' THEN 3
END attribute,
chect_at as check_at
FROM cf_account;

-- Extract data from table "cf_statement" order by transfer_day
-- DROP TABLE IF EXISTS public.cf_statement_x
CREATE TABLE IF NOT EXISTS public.cf_statement_x
(
    id SERIAL,
    period smallint,
    transfer_day date NOT NULL,
    transfer_time time with time zone,
    transfer_at timestamp with time zone,
    inflow numeric NOT NULL DEFAULT 0,
    outflow numeric NOT NULL DEFAULT 0,
    attribute smallint NOT NULL DEFAULT 1,
    amount numeric,
    profit numeric,
    check_at date,
    CONSTRAINT pk_cf_x PRIMARY KEY (id)
);


INSERT INTO public.cf_statement_x (PERIOD, 
transfer_day,
transfer_time,
transfer_at, inflow,  outflow, attribute, check_at)
SELECT PERIOD, 
transfer_day,
transfer_time,
transfer_at, inflow,  outflow, attribute,check_at
FROM public.cf_statement
ORDER BY transfer_day ASC



-- calculate amount in table cf_statement
CREATE OR REPLACE FUNCTION aggr_sum_cf_stmt_amount()
  RETURNS void AS
$$
DECLARE
  cf_flow cursor  for select id,inflow, outflow from cf_statement_x; 
  v_id        INT;
  v_inflow    numeric;
  v_outflow   numeric;
  v_amount    numeric;
BEGIN
    v_amount:=0;
    open cf_flow;
    FETCH cf_flow INTO v_id, v_inflow,v_outflow;
    WHILE FOUND LOOP  
	    v_amount:=v_amount+v_inflow-v_outflow;
		RAISE NOTICE 'id: %, amount: %', v_id, v_amount;
		UPDATE "cf_statement_x" set "amount"=v_amount where "id" = v_id;
		FETCH cf_flow INTO v_id, v_inflow,v_outflow;	
    END LOOP; 
    CLOSE cf_flow;
END;
$$ LANGUAGE plpgsql;
SELECT aggr_sum_cf_stmt_amount();


-- calculate cash flow with occupied days.
CREATE OR REPLACE FUNCTION aggr_cf_fit_amount()
  RETURNS void AS
$$
DECLARE
  cf_flow cursor  for select transfer_day,amount from cf_statement_1; 
  transfer_day   date;
  p_transfer_day date;
  s_transfer_day date;
  -- current day
  c_transfer_day date;
  -- end transfer_day Year_12_31
  e_transfer_day date;
  v_amount     numeric;
  -- accumulate amount 
  a_amount    numeric;
  -- current amount of the day
  c_amount    numeric;
  p_amount    numeric;
  v_flow    numeric;
  -- loop count
  v_cnt     int;
  days      int;
BEGIN
    open cf_flow;
    FETCH cf_flow INTO transfer_day, v_amount;
	p_transfer_day:=transfer_day;
	-- start transfer day
	s_transfer_day=transfer_day;
	a_amount=0;
	p_amount=v_amount;
	v_cnt=0;
	e_transfer_day=(date_trunc('year', s_transfer_day)+interval '1 YEAR')::date-1;
    WHILE FOUND LOOP
		days:=transfer_day-p_transfer_day;
		v_cnt:=v_cnt+1;
		IF days>0 and v_cnt!=0 THEN
		    RAISE NOTICE 'pre: %, day: %, a_amount: %', p_transfer_day, days, a_amount;
			RAISE NOTICE 'p_amount: %, v_amount: %', p_amount, v_amount; 
			
		    FOR d in 0..days-1 loop
			    c_transfer_day=p_transfer_day+d;
			    c_amount:=(p_amount*(d+1)+a_amount)/(c_transfer_day-s_transfer_day+1);
		        --RAISE NOTICE 'n of day: %,  day %, current amount: %',c_transfer_day-s_transfer_day+1, c_transfer_day, c_amount;
		    END LOOP;
			a_amount=a_amount+p_amount*days;
			RAISE NOTICE 'Transfer day: %,  day %',transfer_day, a_amount;
		END IF;
		p_transfer_day=transfer_day;
		p_amount=v_amount;
		FETCH cf_flow INTO  transfer_day, v_amount;
		-- Pay attention to the day
		days=e_transfer_day-p_transfer_day;
		IF days=0 THEN
		    RAISE NOTICE 'End transfer day: %, days %, amount %',p_transfer_day, days, p_amount;
			c_amount=(p_amount*1+a_amount)/(p_transfer_day-s_transfer_day+1);
			RAISE NOTICE 'Current day %, current amount: %',p_transfer_day, c_amount;
		ELSE
		    FOR d in 0..days-1 loop
			    c_transfer_day=p_transfer_day+d;
			    c_amount:=(p_amount*(d+1)+a_amount)/(c_transfer_day-s_transfer_day+1);
		        --RAISE NOTICE 'n of day: %,  day %, current amount: %',c_transfer_day-s_transfer_day+1, c_transfer_day, c_amount;
		    END LOOP;
		    RAISE NOTICE 'ELSE end transfer day: %, days %, amount %',p_transfer_day, days, p_amount;
		END IF;
    END LOOP; 
    CLOSE cf_flow;
	RAISE NOTICE 'cnt: %', v_cnt;
END;
$$ LANGUAGE plpgsql;
SELECT aggr_cf_fit_amount();


-- calculate cash flow with occupied days.
CREATE OR REPLACE FUNCTION insert_assets_into_cf_statement_x()
  RETURNS void AS
$$
DECLARE
  cf_stmt cursor for select transfer_day::date, transfer_time::time with time zone,check_at::date,assets 
                     from hbu.cf_account where check_at is not null; 
  v_transfer_day   date;
  v_check_at      date;
  v_transfer_time time with time zone;
  v_assets     numeric;
  v_cnt        int;
BEGIN
    OPEN cf_stmt;
    FETCH cf_stmt INTO v_transfer_day,v_transfer_time,v_check_at,v_assets;
	IF v_transfer_time IS NULL THEN
	    UPDATE cf_statement_x  SET check_at=v_check_at, assets=v_assets
		    WHERE transfer_day=v_transfer_day AND attribute=1;
	END IF;
	v_cnt=0;
    WHILE FOUND LOOP
	    v_cnt=v_cnt+1;
		RAISE NOTICE 'Cnt: %, Transfer at % %,Check at %, assets: %',v_cnt, v_transfer_day,v_transfer_time, v_check_at, v_assets;
		IF v_transfer_time IS NULL THEN
	        UPDATE cf_statement_x  SET check_at=v_check_at, assets=v_assets
		        WHERE transfer_day=v_transfer_day;
		ELSE
		    UPDATE cf_statement_x  SET check_at=v_check_at, assets=v_assets
		        WHERE transfer_day=v_transfer_day and transfer_time=v_transfer_time;
		END IF;
		FETCH cf_stmt INTO v_transfer_day,v_transfer_time,v_check_at,v_assets;
		
		-- Pay attention to the day
    END LOOP; 
	--RAISE NOTICE 'Cnt: %, Check at %, assets: %',v_cnt, v_check_at, v_assets;
    CLOSE cf_stmt;
END;
$$ LANGUAGE plpgsql;
select insert_assets_into_cf_statement_x();