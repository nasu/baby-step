using System;
namespace QuestionOp
{
    class QuestoinOp
    {
        static void Main()
        {
            var s1 = new Sub(){P1=1,P2=2};
            var s2 = (Sub)null;
            Console.WriteLine("{0}", s1?.P1 ?? 0);
            Console.WriteLine("{0}", s2?.P1 ?? 0);
        }
    }
    class Sub
    {
        public int P1 { get; set; }
        public int P2 { get; set; }
    }
}
