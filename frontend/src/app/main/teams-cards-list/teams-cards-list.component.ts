import { Component, Input } from '@angular/core';
import { TeamsResource } from '../../services/teams.service';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';

@Component({
  selector: 'app-teams-cards-list',
  templateUrl: './teams-cards-list.component.html',
  styleUrls: ['./teams-cards-list.component.scss'],
})
export class TeamsCardsListComponent {
  @Input()
  teams!: TeamsResource[];

  cols = 3;
  rowHeight = '500px';

  handsetPortrait = false;

  constructor(
    // private dialog: MatDialog,
    private responsive: BreakpointObserver
  ) {}

  ngOnInit() {
    this.responsive
      .observe([
        Breakpoints.TabletPortrait,
        Breakpoints.TabletLandscape,
        Breakpoints.HandsetPortrait,
        Breakpoints.HandsetLandscape,
      ])
      .subscribe((result) => {
        this.cols = 3;
        this.rowHeight = '500px';
        this.handsetPortrait = false;

        const breakpoints = result.breakpoints;
        if (breakpoints[Breakpoints.TabletPortrait]) {
          this.cols = 1;
        } else if (breakpoints[Breakpoints.HandsetPortrait]) {
          this.cols = 1;
          this.rowHeight = '430px';
          this.handsetPortrait = true;
        } else if (breakpoints[Breakpoints.HandsetLandscape]) {
          this.cols = 1;
        } else if (breakpoints[Breakpoints.TabletLandscape]) {
          this.cols = 2;
        }
      });
  }
}
