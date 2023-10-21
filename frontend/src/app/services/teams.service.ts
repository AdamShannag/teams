import { Injectable } from '@angular/core';
import { BehaviorSubject, switchMap, throwError } from 'rxjs';

export interface TeamsResource {
  seqNo: number;
  name: string;
  description: string;
  teamIcon: string;
}

export interface TeamsTable {
  header: { key: string; label: string };
  data: [];
}

@Injectable({
  providedIn: 'root',
})
export class TeamsService {
  private subject = new BehaviorSubject<TeamsResource[]>([]);

  public teams$ = this.subject.asObservable();

  constructor() {
    this.subject.next([
      {
        seqNo: 1,
        name: 'test1',
        description: 'test1 description',
        teamIcon:
          'https://icon-library.com/images/team-icon-png/team-icon-png-4.jpg',
      },
      {
        seqNo: 2,
        name: 'test2',
        description: 'test2 description',
        teamIcon:
          'https://icon-library.com/images/team-icon-png/team-icon-png-4.jpg',
      },
      {
        seqNo: 3,
        name: 'test3',
        description: 'test3 description',
        teamIcon:
          'https://icon-library.com/images/team-icon-png/team-icon-png-4.jpg',
      },
      {
        seqNo: 4,
        name: 'test4',
        description: 'test4 description',
        teamIcon:
          'https://icon-library.com/images/team-icon-png/team-icon-png-4.jpg',
      },
      {
        seqNo: 5,
        name: 'test5',
        description: 'test5 description',
        teamIcon:
          'https://icon-library.com/images/team-icon-png/team-icon-png-4.jpg',
      },
      {
        seqNo: 6,
        name: 'test6',
        description: 'test6 description',
        teamIcon:
          'https://icon-library.com/images/team-icon-png/team-icon-png-4.jpg',
      },
    ]);
  }

  fetchAllTeams() {
    return this.teams$;
  }

  getTeamByName(teamName: string) {
    return this.teams$.pipe(
      switchMap((teams) => {
        if (teams.filter((t) => t.name === teamName).length == 0) {
          return throwError(() => new Error('Error emitted by throwError'));
        }
        return teams.filter((t) => t.name === teamName);
      })
    );
  }
}
