using System;
using System.Collections.Generic;
namespace Solution {
    class Solution {
        static void Main(string[] args) {
		//Input
            var line1 = System.Console.ReadLine().Trim();
            var Needs = Int32.Parse(line1);
            var line2 = System.Console.ReadLine().Trim();
            var N = Int32.Parse(line2);
            int[] flag= new int[N];
            for(var i=0;i<N;i++){flag[i]=0;}
            //flag[N-1]=1;
            int cost=500000000;
            List<string[]> data = new List<string[]>();
            for (var i = 0; i < N; i++) {
                string[] stArrayData = System.Console.ReadLine().Trim().Split(' ');
                data.Add(stArrayData);
            }
            
            // Calculate
            int tempcost=0;
            int tempninku=0;
            for(var i=0; i < Math.Pow(2,N);i++){
                for(var j=0;j<N;j++){
                    flag[j]=(flag[j]==0)?1:0;
                    if(flag[j]==1){break;}
              }
                for(var j=0;j<N;j++){
                    if(flag[j]==1){
                        string[] datum=data[j];
                        tempninku+=Int32.Parse(datum[0]);
                        tempcost+=Int32.Parse(datum[1]);
                    }
                   
                }
                if(tempninku>=Needs&&tempcost<cost){cost=tempcost;}
                tempcost=0;
                tempninku=0;
            }
            System.Console.WriteLine(cost);
        }
    }
}
