import { Component, OnInit, inject } from '@angular/core';
import { MessagesService } from './messages.service';
import {
  MAT_SNACK_BAR_DATA,
  MatSnackBar,
  MatSnackBarRef,
} from '@angular/material/snack-bar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';

export interface SnackBarMessage {
  message: string;
  action: string;
  status: SnackBarStatus;
}

export enum SnackBarStatus {
  SUCCESS = 'success',
  WARN = 'warn',
  FAIL = 'fail',
}

@Component({
  selector: 'messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.css'],
})
export class MessagesComponent implements OnInit {
  durationInSeconds = 2;

  constructor(
    public messagesService: MessagesService,
    private snackBar: MatSnackBar
  ) {}

  ngOnInit() {
    this.messagesService.errors$.subscribe((errors) => {
      errors.forEach((error) => {
        this.openSnackBar(error);
      });
    });
  }

  openSnackBar({
    message,
    action,
    status,
  }: {
    message: string;
    action: string;
    status: SnackBarStatus;
  }) {
    this.snackBar.openFromComponent(SnackBarErrorComponent, {
      duration: this.durationInSeconds * 1000,
      data: {
        message,
        action,
        status,
      },
    });
  }
}

@Component({
  selector: 'snack-bar-component',
  templateUrl: './snack-bar-component.html',
  styles: [
    `
      .bar {
        display: flex;
        justify-content: space-between;
        align-items: center;
      }

      .message-success {
        color: #79de79;
      }

      .message-warn {
        color: #fcfc99;
      }

      .message-fail {
        color: #fb6962;
      }
    `,
  ],
  standalone: true,
  imports: [MatIconModule, MatButtonModule],
})
export class SnackBarErrorComponent implements OnInit {
  snackBarRef = inject(MatSnackBarRef);
  data: SnackBarMessage = inject(MAT_SNACK_BAR_DATA);

  ngOnInit(): void {}
}
