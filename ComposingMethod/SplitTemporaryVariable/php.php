//Problem
function discount($inputVal, $quantity) {
  if ($quantity > 50) {
    $inputVal -= 2;
  }
  ...

//Solution
function discount($inputVal, $quantity) {
  $result = $inputVal;
  if ($quantity > 50) {
    $result -= 2;
  }
  ...