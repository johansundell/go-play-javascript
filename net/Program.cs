﻿ using System;

namespace net
{
    class Program
    {
        static double term(int n)
        {
            double i = n * 2 + 1;
            if( n % 2 == 0 )
                return -4 / ((i - 1) * i * (i + 1));
            else
                return 4 / ((i - 1) * i * (i + 1));
        }

        static double GetPi(int n)
        {
            double f = 3.0;
            for( int i = 1; i <= n; i++ )
                f += term(i);
            return f;
        }

        static void Main(string[] args)
        {
            Console.WriteLine(GetPi(1000 * 1000 * 1000) );
        }
    }
}
