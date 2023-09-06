import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';
import { filter } from 'rxjs/operators';
import { SnackBarMessage } from './messages.component';

@Injectable({
  providedIn: 'root',
})
export class MessagesService {
  private subject = new BehaviorSubject<SnackBarMessage[]>([]);

  errors$: Observable<SnackBarMessage[]> = this.subject
    .asObservable()
    .pipe(filter((msgs) => msgs && msgs.length > 0));

  showErrors(...errors: SnackBarMessage[]) {
    this.subject.next(errors);
  }
}
