import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CargaJSONComponent } from './carga-json.component';

describe('CargaJSONComponent', () => {
  let component: CargaJSONComponent;
  let fixture: ComponentFixture<CargaJSONComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CargaJSONComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CargaJSONComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
