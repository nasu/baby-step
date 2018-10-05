using System;

namespace cli
{
    class Echo
    {
        static void Main(string[] args)
        {
            Console.WriteLine(string.Join(" ", args));
        }
    }
}
