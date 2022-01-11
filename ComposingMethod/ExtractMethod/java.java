//This can be extract into method. So whenever we need to reuse the code.
void printOwing() {
    printBanner();
    System.out.Println("name: "+name);
    System.out.Println("amount: "+getOut);
}

//This is how extract method work
void printOwing() {
    printBanner();
    printDetails(getOutstanding());
}
void printDetails(double outstanding) {
    System.out.Println("name"+name);
    System.out.Println("amount"+outstanding);
}