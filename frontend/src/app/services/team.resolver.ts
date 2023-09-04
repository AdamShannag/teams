import { ResolveFn, Router } from '@angular/router';
import { TeamsResource, TeamsService } from './teams.service';
import { Observable, tap, map, catchError, EMPTY } from 'rxjs';
import { inject } from '@angular/core';

export const teamResolver: ResolveFn<TeamsResource> = (
  route,
  state
): Observable<TeamsResource> => {
  const teamService = inject(TeamsService);
  const router = inject(Router);

  const teamName = route.paramMap.get('teamName');

  return teamService.getTeamByName(teamName!).pipe(
    catchError(() => {
      router.navigateByUrl('/main/not-found');
      return EMPTY;
    })
  );
};
