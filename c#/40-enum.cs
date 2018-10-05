using System;

namespace Foo
{
    enum Animal
    {
        Cat,
        [Name("いぬ")]
        Dog,
        [Name("ねずみ")]
        Rat,
    }
    [AttributeUsage( AttributeTargets.Field )]
    static class AnimalExtension
    {
        public static string Name(Animal a)
        {
            string[] names = { "ねこ", "いぬ", "ねずみ" };
            return names[(int)a];
        }
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
