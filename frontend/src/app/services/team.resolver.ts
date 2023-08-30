import { ResolveFn } from '@angular/router';
import { TeamsResource, TeamsService } from './teams.service';
import { Observable } from 'rxjs';
import { inject } from '@angular/core';

export const teamResolver: ResolveFn<TeamsResource> = (
  route,
  state
): Observable<TeamsResource> => {
  const teamService = inject(TeamsService);

  const teamName = route.paramMap.get('teamName');
  return teamService.getTeamByName(teamName!);
};
