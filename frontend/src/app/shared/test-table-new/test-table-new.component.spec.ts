import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TestTableNewComponent } from './test-table-new.component';

describe('TestTableNewComponent', () => {
  let component: TestTableNewComponent;
  let fixture: ComponentFixture<TestTableNewComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [TestTableNewComponent]
    });
    fixture = TestBed.createComponent(TestTableNewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
