import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PasswordReminderComponent } from './password-reminder.component';

describe('PasswordReminderComponent', () => {
  let component: PasswordReminderComponent;
  let fixture: ComponentFixture<PasswordReminderComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [PasswordReminderComponent]
    });
    fixture = TestBed.createComponent(PasswordReminderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
