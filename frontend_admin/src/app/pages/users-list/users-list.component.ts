import { Component } from '@angular/core';
import { FilterUser, User, CsvData } from 'src/app/interfaces/user';
import { UserService } from 'src/app/repositories/user.service';
import { NzMessageService } from 'ng-zorro-antd/message';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { NzModalService, NzModalRef } from 'ng-zorro-antd/modal';
import { UserCreateComponent } from 'src/app/components/user-create/user-create.component';
import { Department } from 'src/app/interfaces/department';
import { Team } from 'src/app/interfaces/team';
import { DepartmentService } from 'src/app/repositories/department.service';
import { TeamService } from 'src/app/repositories/team.service';
import { NzUploadFile } from 'ng-zorro-antd/upload';
import { Observable, Observer } from 'rxjs';
import { CsvTableComponent } from 'src/app/components/csv-table/csv-table.component';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';
import { AuthService } from 'src/app/services/auth.service';
import { environment } from '../../../environments/environment';

@Component({
  selector: 'app-users-list',
  templateUrl: './users-list.component.html',
  styleUrls: ['./users-list.component.scss']
})
export class UsersListComponent {
  loading = false;
  teamOptions: { nzLabel: string; nzValue: string }[] = [];
  editForm !: FormGroup
  filterForm !: FormGroup
  userData: User[] = [];
  teamData: Team[] = [];
  csvData: any[] = [];
  departmentData: Department[] = [];
  clickedId!: string | '';
  totalUsers: number = 0;
  pageIndex!: number;
  pageSize!: number;
  filterd = false;
  accessToken = localStorage.getItem('access_token');
  loginUserId = localStorage.getItem('login_user_id');
  frontendUrl = environment.frontendUrl;
  options: string[] = [];
  data: any[] = [];

  constructor(
    private userSvc: UserService,
    private messageSvc: NzMessageService,
    private fb: FormBuilder,
    private modalSvc: NzModalService,
    private departmentSvc: DepartmentService,
    private teamSvc: TeamService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService,
    private authSvc: AuthService
  ) {
    this.editForm = this.fb.group({
      staff_id: ['', [Validators.required, Validators.pattern('^E[0-9]{5}$')]],
      name: ['', Validators.required],
      display_name: ['', [Validators.required]],
      email: ['', [Validators.required, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]],
      role: ['', [Validators.required]],
      phone: ['', [Validators.pattern('^[0-9]{11}$')]],
      address: [''],
      department_id: ['', [Validators.required]],
      team_id: ['', [Validators.required]]
    });
    this.filterForm = this.fb.group({
      name: [''],
      email: [''],
      department_id: [[]],
      team_id: [[]]
    })
  }

  ngOnInit(): void {
    this.pageIndex = 1;
    this.pageSize = 10;
    this.getUsersLists(this.pageIndex, this.pageSize)
    this.getTeamList()
    this.getDepartmentList()
    this.getUserCount()
    this.editForm.controls['department_id'].valueChanges.subscribe(data => {
      this.editForm.controls['team_id'].setValue('');
    })
  }

  getUserCount() {
    if (this.filterForm.value) {
      this.userSvc.getUserCount(this.filterForm.value).subscribe({
        next: (count) => {
          this.totalUsers = count.count
        }
      })
    }
    else {
      this.userSvc.getUserCount().subscribe({
        next: (count) => {
          this.totalUsers = count.count
        }
      })
    }
  }

  getUsersLists(page: number, limit: number) {
    this.loaderSvc.call();
    this.userSvc.getUsersList(page, limit).subscribe({
      next: (userListData: User[]) => {
        if (!userListData) {
          this.loaderSvc.dissAll();
        }
        this.userData = userListData;
        this.updateDepartmentAndTeamNames();
        this.loaderSvc.dissAll();
      },
      error: (error) => {
        this.loaderSvc.dissAll();
        this.userSvc.handleErrorType(error)
      }
    })
  }

  getTeamList(): void {
    this.teamSvc.getTeamList().subscribe({
      next: (teamListData: Team[]) => {
        this.teamData = teamListData;
        if (teamListData) {
          this.updateDepartmentAndTeamNames();
        }
      },
      error: (error) => {
        console.log(error, "error....")
      }
    })
  }

  getDepartmentList(): void {
    this.departmentSvc.getDepartmentList().subscribe({
      next: (departmentListData: Department[]) => {
        this.departmentData = departmentListData;
        if (departmentListData) {
          this.updateDepartmentAndTeamNames();
        }
      },
      error: (error) => {
        console.log(error, "error....")
      }
    });
  }

  updateDepartmentAndTeamNames() {
    if (this.teamData.length > 0 && this.departmentData.length > 0) {
      if (this.userData) {
        this.userData.forEach((user: User) => {
          const department = this.departmentData.find((dept: Department) => dept._id === user.department_id);
          const team = this.teamData.find((team: Team) => team._id === user.team_id);
          user.department_name = department ? department.name : "";
          user.team_name = team ? team.name : "";
        });
      }
      if (this.csvData) {
        this.csvData.forEach((csv: any) => {
          const department = this.departmentData.find((dept: Department) => dept._id === csv.department_name);
          const team = this.teamData.find((team: Team) => team._id === csv.team_name);
          csv.department_name = department ? department.name : "";
          csv.team_name = team ? team.name : "";

          switch (csv.role) {
            case "manager":
              csv.role = 1
              break;
            case "bse":
              csv.role = 2
              break;
            case "leader":
              csv.role = 3
              break;
            case "sub leader":
              csv.role = 4;
              break;
            case "senior":
              csv.role = 5;
              break;
            case "junior":
              csv.role = 6;
              break
            default:
              csv.role = 6;
          }
        });
      }
    }
  }

  isManagerOrNotManager(id: string): void {
    this.userSvc.getUser(id).subscribe({
      next: (userData: User) => {
        console.log(userData)
      }
    })
  }

  cancel(): void {
    this.messageSvc.info('cancelled');
  }

  confirm(id: string): void {
    this.loaderSvc.call();
    this.userSvc.deleteUser(id).subscribe({
      next: () => {
        this.messageSvc.info(`user deleted successfully`);
        this.getUserCount()
        this.getUsersLists(this.pageIndex, this.pageSize)
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        console.log(error, "error");
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error);
        this.apiSvc.handleMessageError(error);
      }
    })
  }

  startEdit(id: string) {
    this.clickedId = id;
    const selectedData = this.userData.find(data => data._id === id);
    this.editForm.patchValue({
      staff_id: selectedData?.staff_id,
      name: selectedData?.name,
      display_name: selectedData?.display_name,
      email: selectedData?.email,
      role: selectedData?.role,
      phone: selectedData?.phone,
      address: selectedData?.address,
      department_id: selectedData?.department_id,
      team_id: selectedData?.team_id
    });
  }

  cancelEdit(): void {
    this.clickedId = '';
    this.messageSvc.info('cancelled');
  }

  updateUser(id: string) {
    this.loaderSvc.call();
    this.userSvc.updateUser(id, this.editForm.value).subscribe({
      next: (data) => {
        this.authSvc.username.next(data.name);
        this.getUsersLists(this.pageIndex, this.pageSize)
        this.messageSvc.info(`user "${data.display_name}" updated successfully`)
        this.clickedId = '';
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        this.loaderSvc.dismiss();
        this.apiSvc.handleErrorType(error);
        this.apiSvc.handleMessageError(error);
      }
    })
  }

  pageIndexChanged(pageIndex: number): void {
    this.pageIndex = pageIndex;
    if (this.filterd) {
      this.filterUser(this.pageIndex, this.pageSize)
    }
    else {
      this.getUsersLists(this.pageIndex, this.pageSize)
    }
  }

  createUserModal() {
    const modalRef = this.modalSvc.create({
      nzContent: UserCreateComponent,
      nzFooter: null,
      nzClosable: true,
      nzCentered: true,
      nzStyle: { 'width': '550px' },
      nzNoAnimation: true,
      nzMaskClosable: false,
      nzComponentParams: {
        departmentData: this.departmentData,
        teamData: this.teamData
      }
    });
    if (this.userData) {
      modalRef?.componentInstance?.userCreated.subscribe((userData: User) => {
        this.userData = [...this.userData, userData];
        this.getUserCount();
        this.updateDepartmentAndTeamNames();
      });
    }
  }

  filterUser(page: number, index: number): void {
    this.filterd = true;
    this.loaderSvc.call();
    this.userSvc.filterUser(this.filterForm.value, page, index).subscribe({
      next: (filterdUser: User[]) => {
        this.userData = filterdUser
        if (filterdUser) {
          this.updateDepartmentAndTeamNames();
          this.pageIndex = 1;
        }
        this.pageIndex = page;
        this.loaderSvc.dissAll();
      },
      error: (error) => {
        this.apiSvc.handleErrorType(error);
      }
    })
  }

  clearFilterData() {
    this.filterForm.controls['team_id'].setValue([]);
    this.filterForm.controls['department_id'].setValue([]);
    this.filterForm.controls['name'].setValue('');
    this.filterForm.controls['email'].setValue('');
  }

  transformFile = (file: NzUploadFile): Observable<Blob> =>
    new Observable((observer: Observer<Blob>) => {
      const fileName = file.name;
      const fileExtension = fileName.slice(fileName.lastIndexOf('.') + 1).toLowerCase();
      const isCsv = fileExtension === 'csv';
      if (!isCsv) {
        this.messageSvc.error('You can only upload CSV file!')
      }
      else {
        const reader = new FileReader();
        reader.readAsText(file as any);
        reader.onload = () => {
          const csvData = reader.result as string;
          const csvRecordsArray = csvData.split(/\r\n|\n/).filter(value => value !== '');
          const headers = csvRecordsArray[0].split(',');
          const csvDataResult: any[] = [];

          for (let i = 1; i < csvRecordsArray.length; i++) {
            const record = csvRecordsArray[i].split(',');
            if (record[5] !== "") {
              record[5] = "0" + record[5];
            }
            const csvObject: any = {};

            for (let j = 0; j < headers.length; j++) {
              if (headers[j] === 'role') {
                csvObject[headers[j]] = +record[j];
              } else {
                csvObject[headers[j]] = record[j];
              }
            }
            csvDataResult.push(csvObject);
          }
          const modalRef = this.modalSvc.create({
            nzContent: CsvTableComponent,
            nzFooter: null,
            nzClosable: false,
            nzCentered: true,
            nzStyle: { 'width': '1400px' },
            nzNoAnimation: true,
            nzMaskClosable: false,
            nzComponentParams: {
              csvData: csvDataResult,
              teamData: this.teamData,
              departmentData: this.departmentData,
            }
          });
          if (this.csvData) {
            modalRef?.componentInstance?.csvCreated.subscribe(() => {
              this.getUserCount()
              this.getUsersLists(this.pageIndex, this.pageSize);
            })
          }
        };
      }
    });

  downloadTableDataToCsv(): void {
    const fileName = 'user_data.csv';
    this.userSvc.getUsersList(0, 0).subscribe({
      next: (userListData: User[]) => {
        this.csvData = userListData.map(user => ({
          staff_id: user.staff_id,
          name: user.name,
          display_name: user.display_name,
          email: user.email,
          role: user.role,
          phone: user.phone,
          address: user.address,
          department_name: user.department_id,
          team_name: user.team_id,
        }))
        this.updateDepartmentAndTeamNames();
        this.downloadCsv(this.csvData, fileName)  
      },
      error: (error) => {
        this.userSvc.handleErrorType(error)
      }
    })
  }

  downloadCsv(data: CsvData[], fileName: string): void {
    const csvContent = this.convertToCsv(data);
    this.saveToFile(csvContent, fileName);
  }

  private convertToCsv(data: CsvData[]): string {
    const header = Object.keys(data[0]).join(',');
    const rows = data.map(item => Object.values(item).join(','));
    return `${header}\n${rows.join('\n')}`;
  }

  private saveToFile(content: string, fileName: string): void {
    const blob = new Blob([content], { type: 'text/csv' });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = fileName;
    a.click();
    window.URL.revokeObjectURL(url);
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
