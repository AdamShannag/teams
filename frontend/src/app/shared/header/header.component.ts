import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { MessagesService } from '../messages/messages.service';
import {
  SnackBarMessage,
  SnackBarStatus,
} from '../messages/messages.component';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss'],
})
export class HeaderComponent {
  snackBarStatus: typeof SnackBarStatus = SnackBarStatus;
  constructor(
    public auth: AuthService,
    private messagesService: MessagesService
  ) {}

  throwMessage(message: SnackBarMessage) {
    this.messagesService.showErrors(message);
  }
}
