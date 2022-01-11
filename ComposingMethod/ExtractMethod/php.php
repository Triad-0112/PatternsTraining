//This can be extract into method. So whenever we need to reuse the code.
function printOwing(){
    $this->printBanner();
    print("name:   ". $this->name);
    print("amount: ". $this->getOutstanding())
}

//This is how extract method work
function printOwing(){{
    $this->printBanner();
    $this->printDetails($this->getOutstanding());
}
function printDetails($outstanding) {
    print("name:   ". $this->name);
    print("amount: ". $this->getOutstanding())
}