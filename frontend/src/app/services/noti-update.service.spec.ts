import { TestBed } from '@angular/core/testing';

import { NotiUpdateService } from './noti-update.service';

describe('NotiUpdateService', () => {
  let service: NotiUpdateService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(NotiUpdateService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
