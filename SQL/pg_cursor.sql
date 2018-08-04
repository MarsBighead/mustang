CREATE OR REPLACE FUNCTION c_student_to_person()
  RETURNS void AS
$$
DECLARE
  c_stu cursor  for select student_no,student_name from student  ; 
  v_no    INT;
  v_name varchar(40);
BEGIN
    open c_stu;
    FETCH c_stu INTO v_no, v_name;
    while found loop  
          INSERT INTO person(student_no,student_name) VALUES (v_no,v_name);
          FETCH c_stu INTO v_no, v_name;
    end loop; 
  
    CLOSE c_stu;
end;  
$$ LANGUAGE plpgsql;
