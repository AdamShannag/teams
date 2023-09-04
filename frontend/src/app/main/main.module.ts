import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home/home.component';
import { MainRoutingModule } from './main-routing.module';
import { TeamsCardsListComponent } from './teams-cards-list/teams-cards-list.component';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatButtonModule } from '@angular/material/button';
import { TeamComponent } from './team/team.component';
import { TeamNotFoundComponent } from './team-not-found/team-not-found.component';

@NgModule({
  declarations: [HomeComponent, TeamsCardsListComponent, TeamComponent, TeamNotFoundComponent],
  imports: [
    CommonModule,
    MainRoutingModule,
    MatIconModule,
    MatCardModule,
    MatGridListModule,
    MatButtonModule,
  ],
})
export class MainModule {}
