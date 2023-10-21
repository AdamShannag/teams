import { inject } from '@angular/core';
import { ResolveFn, Router } from '@angular/router';
import { Observable, BehaviorSubject, catchError, EMPTY } from 'rxjs';

export const tableResolver: ResolveFn<any> = (
  route,
  state
): Observable<any> => {
  const router = inject(Router);

  return mockObservable().pipe(
    catchError(() => {
      router.navigateByUrl('/error-page');
      return EMPTY;
    })
  );
};

const mockObservable = (): Observable<any> => {
  let subject = new BehaviorSubject<any>(null);
  let tableData$ = subject.asObservable();

  let headers = [
    { key: 'seqNo', label: '#' },
    { key: 'name', label: 'Name' },
    { key: 'description', label: 'Description' },
  ];

  let data: any[] = [];

  subject.next({ headers, data });

  return tableData$;
};
