#include<stdio.h>

void swap(int *i,int *j){
    int t;
    t = *i;
    *i = *j;
    *j = t;
}

int main(){
    int i=1,j=2,lg,tmp;
    swap(&i,&j);
    printf("swap i=%d, j=%d\n\n",i,j);

    int a[]={3,7,8,5,12,14,21,15,18,14};
    lg=sizeof(a)/sizeof(a[0]);
    for (i=0;i<lg;i++){
        for (j=i+1;j<lg;j++){
            if (a[j]<a[i]){
                swap(&a[j],&a[i]);
            }
        }
        int x=0;
        for (x=0;x<lg;x++){
            printf("%d ",a[x]);
        }
        printf("\n");
    }
    printf("\n\n#####\n");
    for (i=0;i<lg;i++){
        printf("%d ",a[i]);
    }
    printf("\n");
}
