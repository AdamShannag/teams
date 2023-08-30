import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TeamsCardsListComponent } from './teams-cards-list.component';

describe('TeamsCardsListComponent', () => {
  let component: TeamsCardsListComponent;
  let fixture: ComponentFixture<TeamsCardsListComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TeamsCardsListComponent]
    });
    fixture = TestBed.createComponent(TeamsCardsListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
