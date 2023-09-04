import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { catchError, throwError } from 'rxjs';
import { TeamsResource } from 'src/app/services/teams.service';
import { tap } from 'rxjs/operators';

@Component({
  selector: 'app-team',
  templateUrl: './team.component.html',
  styleUrls: ['./team.component.scss'],
})
export class TeamComponent {
  team: TeamsResource = this.route.snapshot.data['team'];

  constructor(private route: ActivatedRoute) {
    this.route.data.subscribe(({ team }) => {
      this.team = team;
    });
  }

  ngOnInit() {}
}
