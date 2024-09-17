#include <stdio.h>
#include <string.h>

void print(int arr[], int n) {
	int i =0;
	for(i=0; i <n; i++) {
	   printf("%d ", arr[i]);
	}
	printf("\n");
}


void merge(int arr[], int l, int mid, int r) {
//	printf("Inside merge l=%d m=%d r=%d", l, mid, r);
        int n1 = mid -l +1;
	int n2 = r-mid;
	int L[n1];
	int R[n2];
	int i,j,k;

	for(i=0;i <n1;i++) {
		L[i] = arr[l+i];
	}
	for(j=0;j<n2;j++) {
		R[j] = arr[mid+j+1];
	}
   
	i = j =0;
	k =l;
	while(i<n1 && j <n2) {
	  if(L[i] < R[j]) {
	    arr[k] = L[i];
	    i++;
	  } else {
	    arr[k] = R[j];
            j++;
          }
	  k++;
        }
//	printf("i=%d j=%d k=%d", i,j,k);
        while(i<n1) {
	    arr[k] = L[i];
            i++;
	    k++;
        }
        while(j<n2) {
	    arr[k] = R[j];
            j++;
	    k++;
        }

}	
void mergeSort(int arr[], int l, int r) {
	if(l <r) {
	    int mid = (l+(r-1))/2;
	    mergeSort(arr, l, mid);
	    mergeSort(arr, mid+1, r);
	    merge(arr, l, mid, r);
	}
}	
int main() {
	int arr[] = {2,-3,1,3,5,1};
	printf("\n%ld\n", sizeof(arr)/sizeof(int));
	int n = sizeof(arr)/sizeof(int);
	print(arr,n);
	mergeSort(arr,0, n-1);
	print(arr, n);
	return 0;
}
