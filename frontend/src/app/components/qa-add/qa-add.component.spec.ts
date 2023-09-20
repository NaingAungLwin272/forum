import { ComponentFixture, TestBed } from '@angular/core/testing';

import { QaAddComponent } from './qa-add.component';

describe('QaAddComponent', () => {
  let component: QaAddComponent;
  let fixture: ComponentFixture<QaAddComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [QaAddComponent]
    });
    fixture = TestBed.createComponent(QaAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
