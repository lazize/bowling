using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Game;
using System.Text;

namespace BowlingTest
{
    [TestClass]
    public class UnitTest1
    {
        [TestMethod]
        public void TestBownlingUnder10()
        {
            for (int i = 0; i < 10; i++)
            {
                for (int j = 0; j < (10 -i); j++)
                {
                    int shouldBeValue = 0;
                    StringBuilder shouldBeString = new StringBuilder();
                    Bowling game = new Bowling();

                    for (int k = 0; k < 10; k++)
                    {
                        shouldBeValue += i + j;
                        shouldBeString.Append(i);
                        shouldBeString.Append(j);
                        shouldBeString.Append(",");
                        game.AddFrame(i, j);
                    }
                    // Remove last ","
                    shouldBeString.Remove(shouldBeString.Length - 1, 1);

                    int score = game.CalculateScore();
                    Assert.AreEqual<int>(shouldBeValue, score);

                    string visual = game.ToString();
                    Assert.AreEqual<string>(shouldBeString.ToString(), visual);
                }
            }
        }
        [TestMethod]
        public void TestBownling11()
        {
            int shouldBeValue = 11;
            string shouldBeString = "00,00,00,00,00,00,00,00,00,1/1";

            Bowling game = new Bowling();
            game.AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddFrameExtra(1, 9, 1);

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
        [TestMethod]
        public void TestBownling20()
        {
            int shouldBeValue = 20;
            string shouldBeString = "00,00,00,00,00,00,00,00,00,X1/";

            Bowling game = new Bowling();
            game.AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddMiss().AddFrameExtra(10, 1, 9);

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
        [TestMethod]
        public void TestBownling270()
        {
            int shouldBeValue = 270;
            string shouldBeString = "X,X,X,X,X,X,X,X,X,9/1";

            Bowling game = new Bowling();
            game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrameExtra(9, 1, 1);

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
        [TestMethod]
        public void TestBownling271()
        {
            int shouldBeValue = 271;
            string shouldBeString = "X,X,X,X,X,X,X,X,X,1/X";

            Bowling game = new Bowling();
            game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrameExtra(1, 9, 10);

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
        [TestMethod]
        public void TestBownling300()
        {
            int shouldBeValue = 300;
            string shouldBeString = "X,X,X,X,X,X,X,X,X,XXX";

            Bowling game = new Bowling();
            game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrameExtra(10, 10, 10);

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
        [TestMethod]
        public void TestBownling10AsStrike()
        {
            int shouldBeValue = 10;
            string shouldBeString = "X";

            Bowling game = new Bowling();
            game.AddStrike();

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
        [TestMethod]
        public void TestBownling10AsSpare()
        {
            int shouldBeValue = 10;
            string shouldBeString = "5/";

            Bowling game = new Bowling();
            game.AddSpare(5);

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
        [TestMethod]
        public void TestBownling245()
        {
            int shouldBeValue = 245;
            string shouldBeString = "X,X,X,X,X,X,X,X,X,11";

            Bowling game = new Bowling();
            game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrame(1, 1);

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
        [TestMethod]
        public void TestBownling273()
        {
            int shouldBeValue = 273;
            string shouldBeString = "X,X,X,X,X,X,X,X,X,X11";

            Bowling game = new Bowling();
            game.AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddStrike().AddFrameExtra(10, 1, 1);

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
        [TestMethod]
        public void TestBownling15()
        {
            int shouldBeValue = 15;
            string shouldBeString = "11,1/,11";

            Bowling game = new Bowling();
            game.AddFrame(1, 1).AddSpare(1).AddFrame(1, 1);

            int score = game.CalculateScore();
            Assert.AreEqual<int>(shouldBeValue, score);

            string visual = game.ToString();
            Assert.AreEqual<string>(shouldBeString, visual);
        }
    }
}
