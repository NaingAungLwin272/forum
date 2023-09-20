import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Department } from 'src/app/interfaces/department';
import { CsvData, User } from 'src/app/interfaces/user';
import { NzMessageService } from 'ng-zorro-antd/message';
import { AbstractControl, FormArray, FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Team } from 'src/app/interfaces/team';
import { UserService } from 'src/app/repositories/user.service';
import { NzModalRef } from 'ng-zorro-antd/modal';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-csv-table',
  templateUrl: './csv-table.component.html',
  styleUrls: ['./csv-table.component.scss']
})
export class CsvTableComponent {
  @Input() departmentData: Department[] = [];
  @Input() csvData: User[] = [];
  @Input() teamData: Team[] = [];
  @Output() csvCreated: EventEmitter<any> = new EventEmitter<any>()
  departmentName!: string;
  teamName!: string;
  csvForm!: FormGroup;
  options: string[] = [];

  constructor(
    private messageSvc: NzMessageService,
    private fb: FormBuilder,
    private userSvc: UserService,
    private modal: NzModalRef,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) { }

  ngOnInit() {
    this.csvForm = this.fb.group({
      user: this.fb.array([])
    });
    this.csvData.map((user: User) => {
      if (this.departmentData && this.teamData) {
        const department = this.departmentData.find((dept: Department) => dept.name === user.department_name);
        const team = this.teamData.find((team: Team) => team.name === user.team_name);
        department ? user.department_id = department._id : user.department_id = '';
        const checkTeam = this.teamData.some(e => e.department_id === department?._id && e._id === team?._id);
        team && checkTeam ? user.team_id = team._id : user.team_id = '';
        const control = <FormArray>this.csvForm.controls['user'];
        const edit = department && team && checkTeam ? true : false;
        // push user form
        control.push(this.getUserFormGroup(user, edit));
      } else {
        const control = <FormArray>this.csvForm.controls['user'];
        control.push(this.getUserFormGroup(user, false));
      }
    });
  }

  private getUserFormGroup(user: User, edit: boolean) {
    return this.fb.group({
      staff_id: [user.staff_id, [Validators.required, Validators.pattern('^E[0-9]{5}$')]],
      name: [user.name, [Validators.required]],
      email: [user.email, [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]],
      display_name: [user.display_name, [Validators.required]],
      role: [user.role, [Validators.required]],
      department_id: [user.department_id, [Validators.required]],
      team_id: [user.team_id, [Validators.required]],
      phone: [user.phone, [Validators.pattern('^[0-9]{11}$')]],
      address: [user.address],
      dob: [null],
      edit: [edit]
    });
  }

  getUserControls() {
    return (this.csvForm.get('user') as FormArray).controls;
  }

  clearDepartment(index: number) {
    const control = this.csvForm.get('user') as FormArray;
    const userFrom = control.at(index);
    userFrom.get('team_id')?.setValue('');
  }

  confirmDelete(i: number): void {
    const control = <FormArray>this.csvForm.controls['user'];
    control.removeAt(i);
    this.updateIndexes();
  }

  updateIndexes(): void {
    const control = this.csvForm.get('user') as FormArray;
    control.controls.forEach((controlItem) => {
      controlItem.get('edit')?.setValue(true);
    });
  }

  cancel() {
    this.messageSvc.info("Cancelled");
  }

  startEdit(index: number) {
    const control = this.csvForm.get('user') as FormArray;
    const userFrom = control.at(index);
    if (userFrom) {
      userFrom.get('edit')?.setValue(false);
    }
  }

  cancelEdit(index: number): void {
    const control = this.csvForm.get('user') as FormArray;
    const userFrom = control.at(index);
    if (userFrom) {
      userFrom.get('edit')?.setValue(true);
    }
    this.messageSvc.info('Cancelled');
  }

  updateUser(userFrom: AbstractControl) {
    if (userFrom.valid) {
      userFrom.get('edit')?.setValue(true);
    }
  }

  createUserWithCsv(): void {
    this.loaderSvc.call();
    const userArray: User[] = this.csvForm.value.user;
    const emailSet = new Set<string>();
    const displayNameSet = new Set<string>();
    const staffIdSet = new Set<string>();

    let hasEmailDuplicates = false;
    let hasDisplayNameDuplicates = false;
    let hasStaffIdDuplicates = false;

    for (const user of userArray) {
      if (emailSet.has(user.email)) {
        hasEmailDuplicates = true;
      }
      if (displayNameSet.has(user.display_name)) {
        hasDisplayNameDuplicates = true;
      }
      if (staffIdSet.has(user.staff_id)) {
        hasStaffIdDuplicates = true;
      }

      emailSet.add(user.email);
      displayNameSet.add(user.display_name);
      staffIdSet.add(user.staff_id);
    }

    if (hasEmailDuplicates) {
      this.loaderSvc.dismiss();
      this.messageSvc.error(
        "email is duplicate"
      )
      return;
    }

    if (hasDisplayNameDuplicates) {
      this.loaderSvc.dismiss();
      this.messageSvc.error(
        "display name is duplicate"
      )
      return;
    }

    if (hasStaffIdDuplicates) {
      this.messageSvc.error(
        "staff id is duplicated"
      )
      this.loaderSvc.dismiss();
      return;
    }

    this.userSvc.createUserWithCsv(userArray).subscribe({
      next: (data: User[]) => {
        this.modal.destroy();
        this.csvCreated.emit(data);
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        if (error.status === 401) {
          this.loaderSvc.dissAll();
        }
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error);
        this.apiSvc.handleMessageError(error);
      }
    });
  }

  cancelCreate() {
    this.modal.close();
  }

  getDepartName(id: string): string {
    return this.departmentData?.find((department: Department) => department._id === id)?.name || '';
  }

  getTeamName(id: string): string {
    return this.teamData?.find((team: Team) => team._id === id)?.name || '';
  }

  onInput(e: Event): void {
    const value = (e.target as HTMLInputElement).value;
    if (!value || value.indexOf('@') >= 0) {
      this.options = [];
    } else {
      this.options = ['gmail.com', 'metateammyanmar.com'].map(domain => `${value}@${domain}`);
    }
  }
}
