import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { TeamComponent } from './team/team.component';
import { teamResolver } from '../services/team.resolver';
import { TeamNotFoundComponent } from './team-not-found/team-not-found.component';

const routes: Routes = [
  {
    path: '',
    redirectTo: 'teams',
    pathMatch: 'full',
  },
  {
    path: 'teams',
    component: HomeComponent,
  },
  {
    path: 'teams/:teamName',
    component: TeamComponent,
    resolve: {
      team: teamResolver,
    },
  },
  {
    path: 'not-found',
    component: TeamNotFoundComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class MainRoutingModule {}
