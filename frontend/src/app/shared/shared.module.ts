import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatMenuModule } from '@angular/material/menu';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { MatTabsModule } from '@angular/material/tabs';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatToolbarModule } from '@angular/material/toolbar';
import { RouterModule } from '@angular/router';
import { MatGridListModule } from '@angular/material/grid-list';
import { SideNavItemComponent } from './side-nav-item/side-nav-item.component';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { LoadingComponent } from './loading/loading.component';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MessagesComponent } from './messages/messages.component';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatRadioModule } from '@angular/material/radio';
import { ReactiveFormsModule } from '@angular/forms';
import { NavigatorComponent } from './navigator/navigator.component';
import { DashboardComponent } from './dashboard/dashboard.component';

@NgModule({
  declarations: [
    SideNavItemComponent,
    LoadingComponent,
    MessagesComponent,
    NavigatorComponent,
    DashboardComponent,
  ],
  imports: [
    RouterModule,
    CommonModule,
    MatMenuModule,
    MatButtonModule,
    MatIconModule,
    MatCardModule,
    MatTabsModule,
    MatSidenavModule,
    MatListModule,
    MatToolbarModule,
    MatGridListModule,
    MatProgressSpinnerModule,
    MatSnackBarModule,
    MatInputModule,
    MatSelectModule,
    MatRadioModule,
    ReactiveFormsModule,
  ],
  exports: [
    MatProgressSpinnerModule,
    LoadingComponent,
    MessagesComponent,
    NavigatorComponent,
    DashboardComponent,
  ],
})
export class SharedModule {}
