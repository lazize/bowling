using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Game
{
    public class Bowling
    {
        public Frame Head { get; private set; }
        public Frame Tail { get; private set; }

        private void AddFrameInternal(Frame frame)
        {
            if (Head == null)
                Head = frame;
            else
                Tail.Next = frame;
            Tail = frame;
        }

        public Bowling AddFrame(int ball1, int ball2)
        {
            AddFrameInternal(new Frame(ball1, ball2));
            return this;
        }
        public Bowling AddFrameExtra(int ball1, int ball2, int extraBall)
        {
            AddFrameInternal(new ExtraBallFrame(ball1, ball2, extraBall));
            return this;
        }

        public Bowling AddStrike()
        {
            AddFrameInternal(new Frame(10, 0));
            return this;
        }
        public Bowling AddSpare(int ball1)
        {
            AddFrameInternal(new Frame(ball1, Frame.FRAME_MAX_PIN - ball1));
            return this;
        }
        public Bowling AddMiss()
        {
            AddFrameInternal(new Frame(0, 0));
            return this;
        }

        public int CalculateScore()
        {
            int score = 0;
            for (var item = Head; item != null; item = item.Next)
            {
                score += item.Value();
            }
            return score;
        }

        public override string ToString()
        {
            StringBuilder builder = new StringBuilder();
            for (var item = Head; item != null; item = item.Next)
            {
                builder.Append(item.ToString());
                builder.Append(",");
            }
            // Remove last ","
            builder.Remove(builder.Length - 1, 1);
            return builder.ToString();
        }
    }

    public class Frame
    {
        public const int FRAME_MAX_PIN = 10;

        public int FirstBall { get; private set; }
        public int SecondBall { get; private set; }
        public bool IsStrike
        {
            get
            {
                return (this.FirstBall == FRAME_MAX_PIN);
            }
        }
        public bool IsSpare
        {
            get
            {
                if (IsStrike)
                    return false;
                return ((this.FirstBall + this.SecondBall) == FRAME_MAX_PIN);
            }
        }

        // Chain of responsibility
        public Frame Next { get; internal set; }

        public Frame(int firstBall, int secondBall)
        {
            this.FirstBall = firstBall;
            this.SecondBall = secondBall;
        }

        public override string ToString()
        {
            if (IsStrike)
                return "X";
            if (IsSpare)
                return $"{FirstBall}/";
            return $"{FirstBall}{SecondBall}";
        }

        public virtual int Value()
        {
            int sum = FirstBall + SecondBall;
            if (Next != null)
            {
                if (IsSpare)
                {
                    sum += Next.FirstBall;
                }
                if (IsStrike)
                {
                    sum += Next.FirstBall + Next.SecondBall;
                    if (Next.Next != null && Next.IsStrike)
                    {
                        sum += Next.Next.FirstBall;
                    }
                }
            }
            return sum;
        }
    }

    public class ExtraBallFrame : Frame
    {
        public int ExtraBall { get; private set; }

        public ExtraBallFrame(int firstBall, int secondBall, int extraBall) : base(firstBall, secondBall)
        {
            ExtraBall = extraBall;
        }

        public override int Value()
        {
            return FirstBall + SecondBall + ExtraBall;
        }

        public override string ToString()
        {
            StringBuilder builder = new StringBuilder();
            builder.Append(FirstBall == Frame.FRAME_MAX_PIN ? "X" : $"{FirstBall}");
            builder.Append(SecondBall == Frame.FRAME_MAX_PIN ? "X" : (FirstBall + SecondBall) == Frame.FRAME_MAX_PIN ? "/" : $"{SecondBall}");
            builder.Append(ExtraBall == Frame.FRAME_MAX_PIN ? "X" : (FirstBall + SecondBall) != Frame.FRAME_MAX_PIN && (SecondBall + ExtraBall) == Frame.FRAME_MAX_PIN ? "/" : $"{ExtraBall}");
            return builder.ToString();
        }
    }
}
