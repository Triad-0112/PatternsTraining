//Problem
function getRating() {
  return ($this->moreThanFiveLateDeliveries()) ? 2 : 1;
}
function moreThanFiveLateDeliveries() {
  return $this->numberOfLateDeliveries > 5;
}

//Solution
function getRating() {
  return ($this->numberOfLateDeliveries > 5) ? 2 : 1;
}