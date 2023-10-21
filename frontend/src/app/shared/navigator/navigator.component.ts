import { Component, inject } from '@angular/core';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { Observable } from 'rxjs';
import { map, shareReplay } from 'rxjs/operators';
import { AuthService } from 'src/app/services/auth.service';
import { MessagesService } from '../messages/messages.service';
import {
  SnackBarStatus,
  SnackBarMessage,
} from '../messages/messages.component';

@Component({
  selector: 'app-navigator',
  templateUrl: './navigator.component.html',
  styleUrls: ['./navigator.component.scss'],
})
export class NavigatorComponent {
  private breakpointObserver = inject(BreakpointObserver);
  private messagesService = inject(MessagesService);

  public auth = inject(AuthService);

  snackBarStatus: typeof SnackBarStatus = SnackBarStatus;

  throwMessage(message: SnackBarMessage) {
    this.messagesService.showErrors(message);
  }

  isHandset$: Observable<boolean> = this.breakpointObserver
    .observe(Breakpoints.Handset)
    .pipe(
      map((result) => result.matches),
      shareReplay()
    );
}
