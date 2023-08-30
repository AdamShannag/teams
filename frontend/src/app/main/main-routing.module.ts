import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { TeamComponent } from './team/team.component';
import { teamResolver } from '../services/team.resolver';

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
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class MainRoutingModule {}
