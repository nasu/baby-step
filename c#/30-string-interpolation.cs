using System;
namespace StringInterPolation
{
    class StringInterPolation
    {
        static void Main()
        {
            int i = 1;
            string s = "foo";
            Console.WriteLine($"{i:00},{s}"); // "01,foo"

            var msg = "{i:00},{s}";
            Console.WriteLine($"{msg}"); // "{i:00},{s}"
            Console.WriteLine($"{$"{msg}"}"); // "{i:00},{s}"
            Console.WriteLine($"{msg}".Replace("{i:00}", $"{i:00}").Replace("{s}", $"{s}")); // "01,foo"
        }
    }
}
