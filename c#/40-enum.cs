using System;

namespace Foo
{
    public enum Animal
    {
        Cat,
        Dog,
        Rat,
    }
    class Foo
    {
        static void Main()
        {
            foreach (var a in Enum.GetValues(typeof(Animal)))
            {
                Console.WriteLine((int)a);
            }
            foreach (var a in Enum.GetNames(typeof(Animal)))
            {
                Console.WriteLine(a);
            }
        }
    }
}
