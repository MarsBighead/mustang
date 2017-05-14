<?php
interface AdvancedArithmetic{
    public function divisorSum($n);
}



//Write your code here
class Calculator implements AdvancedArithmetic{
    public function divisorSum($n){
        $h= $n/2;
        $sum =0;
        for ($i=1; $i<=$h; $i++){
            if ( $n%$i==0 ){
               $sum+=$i;
            }else{
                continue;
            }
        }
        $sum+=$n;
        return $sum;
    }
}

$n=intval(fgets(STDIN));
$myCalculator=new Calculator();
if($myCalculator instanceof AdvancedArithmetic)//checking if Calculator has implemented AdvancedArithemtic
{
    $sum=$myCalculator->divisorSum($n);
    echo "I implemented: AdvancedArithmetic\n".$sum;
}
else
{
    echo "Wrong answer";// You will get this output if you dont implement
}
?>
