#include <iostream>
using namespace std;

long sumOfEvenFibonnaci(long n){
    long a =2, b = 8;
    long sum = 2;
    while (b<=n){
        sum +=b;
        long tmp = a;
        a = b;
        b = 4*b + tmp;
    }
    return sum;
}

int main(){
    int T;
    cin >> T;
    for (size_t i = 0; i < T; i++){
        long n;
        cin >> n;
        cout << sumOfEvenFibonnaci(n)<<endl;
    }
    return 0;
}