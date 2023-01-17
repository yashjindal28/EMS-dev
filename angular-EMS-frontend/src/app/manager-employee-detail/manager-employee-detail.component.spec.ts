import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagerEmployeeDetailComponent } from './manager-employee-detail.component';

describe('ManagerEmployeeDetailComponent', () => {
  let component: ManagerEmployeeDetailComponent;
  let fixture: ComponentFixture<ManagerEmployeeDetailComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagerEmployeeDetailComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ManagerEmployeeDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
