import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SideNavItemComponent } from './side-nav-item.component';

describe('SideNavItemComponent', () => {
  let component: SideNavItemComponent;
  let fixture: ComponentFixture<SideNavItemComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SideNavItemComponent]
    });
    fixture = TestBed.createComponent(SideNavItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
