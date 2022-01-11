//This can be extract into method. So whenever we need to reuse the code.
void printOwing()
{
    this.PrintBanner();
    Console.WriteLine("name:   "+this.name);
    Console.WriteLine("amount: "+this.GetOutstanding());
}

//This is how extract method work
void PrintOwing()
{
    this.PrintBanner();
    this.PrintDetails();
}

void PrintDetails()
{
    Console.WriteLine("name:   "+this.name);
    Console.WriteLine("amount: "+this.GetOutstanding());
}