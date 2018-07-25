using System;
using System.Collections.Generic;
namespace Sort
{
    class Sort
    {
        static void Main()
        {
            var list = new List<Sub>();
            list.Add(new Sub(1,10));
            list.Add(new Sub(5,50));
            list.Add(new Sub(4,40));
            list.Add(new Sub(2,20));
            list.Add(new Sub(3,30));
            list.ForEach(_ => Console.Write("{0},", _.P1));
            Console.Write("\n");

            list.Sort((a,b) => {
                //DESC
                if (a.P1 < b.P1) return 1;
                if (a.P1 > b.P1) return -1;
                return 0;
            });
            list.ForEach(_ => Console.Write("{0},", _.P1));
            Console.Write("\n");

            list.Sort((a,b) => {
                //ASC
                if (a.P1 < b.P1) return -1;
                if (a.P1 > b.P1) return 1;
                return 0;
            });
            list.ForEach(_ => Console.Write("{0},", _.P1));
            Console.Write("\n");
        }
    }
    class Sub
    {
        public int P1 { get; set; }
        public int P2 { get; set; }
        public Sub(int p1, int p2)
        {
            P1 = p1;
            P2 = p2;
        }
    }
}
