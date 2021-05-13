import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReportesmerkleComponent } from './reportesmerkle.component';

describe('ReportesmerkleComponent', () => {
  let component: ReportesmerkleComponent;
  let fixture: ComponentFixture<ReportesmerkleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReportesmerkleComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ReportesmerkleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
