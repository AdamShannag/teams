import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TeamNotFoundComponent } from './team-not-found.component';

describe('TeamNotFoundComponent', () => {
  let component: TeamNotFoundComponent;
  let fixture: ComponentFixture<TeamNotFoundComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TeamNotFoundComponent]
    });
    fixture = TestBed.createComponent(TeamNotFoundComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
