import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TeamsResource } from 'src/app/services/teams.service';

@Component({
  selector: 'app-team',
  templateUrl: './team.component.html',
  styleUrls: ['./team.component.scss'],
})
export class TeamComponent {
  team!: TeamsResource;

  constructor(private route: ActivatedRoute) {}

  ngOnInit() {
    this.team = this.route.snapshot.data['team'];
  }
}
