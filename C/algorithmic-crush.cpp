#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;
 
typedef long long ll;
 
int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */
    ll n, m, l, r, k, sum=0, max=0;
    cin >> n >> m;
    vector<ll> arr(n+1);
    for(int i=1; i <= n; i++)
        arr[i] = 0;
    for(int i=1; i <= m; i++) {
        cin >> l >> r >> k;
        arr[l] += k;
        if(r <= n) arr[r+1] -= k;
    }
    max = arr[1];
    sum = arr[1];
    for(int i=2; i <= n; i++) {
        sum+=arr[i];
        if(max < sum) max = sum;
    }
    cout << max;
    return 0;
}
