import { TestBed } from '@angular/core/testing';

import { LastpostService } from './lastpost.service';

describe('LastpostService', () => {
  let service: LastpostService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LastpostService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
