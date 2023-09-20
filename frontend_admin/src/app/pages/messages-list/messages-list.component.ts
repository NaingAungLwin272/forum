import { Component } from '@angular/core';
import { User } from 'src/app/interfaces/user';
import { UserService } from 'src/app/repositories/user.service';
import { DepartmentService } from 'src/app/repositories/department.service';
import { TeamService } from 'src/app/repositories/team.service';
import { Department } from 'src/app/interfaces/department';
import { Team } from 'src/app/interfaces/team';
import { FormGroup, Validators, FormBuilder } from '@angular/forms';
import { NzMessageService } from "ng-zorro-antd/message"
import { NzModalService } from 'ng-zorro-antd/modal';
import { NotiService } from 'src/app/repositories/noti.service';
import { Noti } from 'src/app/interfaces/noti';
import { isEmpty } from 'lodash';
import { LoaderService } from 'src/app/services/loader.service';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-messages-list',
  templateUrl: './messages-list.component.html',
  styleUrls: ['./messages-list.component.scss']
})
export class MessagesListComponent {
  userData: User[] = [];
  departmentData: Department[] = [];
  teamData: Team[] = [];
  selectedOption: string = 'User';
  msgForm: FormGroup;
  isLoading = false;

  constructor(
    private userSvc: UserService,
    private departmentSvc: DepartmentService,
    private teamSvc: TeamService,
    private notiSvc: NotiService,
    private formBuilder: FormBuilder,
    private messageSvc: NzMessageService,
    private modalnew: NzModalService,
    private loaderSvc: LoaderService,
    private apiSvc: ApiService
  ) {
    this.msgForm = this.formBuilder.group({
      user_id: [[], Validators.required],
      name: ['', Validators.required],
      link: [''],
      description: ['', Validators.required]
    });
  }

  ngOnInit(): void {
    this.loadData();
  }

  loadData(): void {
    this.getUserList();
    this.getTeamList();
  }

  getUserList() {
    this.loaderSvc.call();
    this.userSvc.getUsersList().subscribe({
      next: (userListData: User[]) => {
        this.userData = userListData;
        this.loaderSvc.dismiss();
      },
      error: (error) => {
        this.loaderSvc.dismiss();
        this.userSvc.handleErrorType(error);
      }
    });
  }

  getTeamList() {
    this.teamSvc.getTeamList().subscribe({
      next: (teamListData: Team[]) => {
        this.teamData = teamListData;

        this.departmentSvc.getDepartmentList().subscribe({
          next: (departmentListData: Department[]) => {
            this.departmentData = departmentListData;
            this.teamData.forEach((team: Team) => {
              const matchingDepartment = this.departmentData.find(department => department._id === team.department_id);
              team.department_name = matchingDepartment ? matchingDepartment.name : 'Unknown Department';
            });
          },
          error: (error) => {
            console.log(error, "error.....")
          }
        });
      },
      error: (error) => {
        console.log(error, "error.....")
      }
    });
  }


  onOptionChange(): void {
    this.msgForm.get('user_id')?.setValue([]);
  }

  createNoti() {
    this.isLoading = true;
    if (this.msgForm.valid && this.userData) {
      let userIDs;
      if (this.selectedOption == 'User') {
        userIDs = this.msgForm.value.user_id
      } else if (this.selectedOption == 'Team') {
        const data = this.userData.filter(user => this.msgForm.value.user_id.includes(user.team_id));
        userIDs = data.map(user => user._id);
      } else if (this.selectedOption == 'Department') {
        const departmentData = this.userData.filter(user => this.msgForm.value.user_id.includes(user.department_id));
        userIDs = departmentData.map(user => user._id)
      }
      const notiData = {
        user_id: userIDs,
        description: this.msgForm.value.description,
        name: this.msgForm.value.name,
        link: this.msgForm.value.link,
        status: true,
        type: 6
      };
      this.loaderSvc.call();
      this.notiSvc.createNoti(notiData).subscribe({
        next: (data: Noti) => {
          if (!isEmpty(data)) {
            this.messageSvc.info(`noti "${this.msgForm.value.name}" created successfully`);
            this.isLoading = false;
          } else {
            this.messageSvc.error(`user not exist`);
            this.isLoading = false;
          }
          this.loaderSvc.dismiss();
        },
        error: (error) => {
          this.loaderSvc.dismiss();
          this.apiSvc.handleErrorType(error);
          this.apiSvc.handleMessageError(error);
          this.isLoading = false;
        }
      })
    }
  }

  clear(): void {
    this.msgForm.get('user_id')?.setValue([]);
    this.msgForm.get('name')?.setValue('');
    this.msgForm.get('link')?.setValue('');
    this.msgForm.get('description')?.setValue('');
  }

  showPopconfirm(): void {
    const confirmationMessage = "Are you sure you want to send a message to the selected users?"
    this.modalnew.confirm({
      nzTitle: confirmationMessage,
      nzOkText: 'OK',
      nzCancelText: 'Cancel',
      nzCentered: true,
      nzOnOk: () => {
        this.createNoti(),
          this.isLoading = false;
      }
    });
  }
}
