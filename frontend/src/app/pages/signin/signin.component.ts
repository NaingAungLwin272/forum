import { ActivatedRoute, Router } from '@angular/router';
import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiService } from 'src/app/services/api.service';
import { AuthService } from 'src/app/services/auth.service';
import { LoaderService } from 'src/app/services/loader.service';
import { Meta } from '@angular/platform-browser';
import { NzModalService } from 'ng-zorro-antd/modal';

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.scss'],
  encapsulation: ViewEncapsulation.None
})
export class SigninComponent implements OnInit {
  signinForm: FormGroup;
  passwordVisible = false;
  redirectUrl!: string;
  options: string[] = [];
  constructor(
    private metaService: Meta,
    private authSvc: AuthService,
    private loaderSvc: LoaderService,
    private router: Router,
    private apiSvc: ApiService,
    private modal: NzModalService,
    private route: ActivatedRoute
  ) {
    this.addTag();
    this.authSvc.clearAll();
    this.signinForm = new FormGroup({
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required, Validators.minLength(6)]),
      is_remember_me: new FormControl(false)
    });
    this.route.queryParams.subscribe((params: any) => {
      this.redirectUrl = params.redirectUrl;
    });
  }

  addTag() {
    this.metaService.addTags([
      { property: 'og:description', content: 'Join our MTM Developer Community Forum where programmers, coders, and tech enthusiasts gather to discuss, share insights, and collaborate on the latest trends, tools, and technologies.' },
      { property: 'og:title', content: 'MTM Community Forum - Connect, Learn & Innovate' },
      { property: 'og:image', content: 'https://res.cloudinary.com/dkjvy425b/image/upload/v1692241903/mtm_community_profile/jpbvzrbszismxvgbiz6v.png' },

    ]);
  }

  ngOnInit(): void {
    this.metaService.addTag({ name: 'description', content: "Article Description" });
    this.modal.closeAll();
  }

  signin(): void {
    this.loaderSvc.call();
    this.authSvc.signin(this.signinForm.value).then(
      (value: { is_authenticated: boolean }) => {
        if (value.is_authenticated === true) {
          this.loaderSvc.dismiss();
          if (this.redirectUrl) {
            this.router.navigateByUrl(this.redirectUrl)
          } else {
            this.router.navigateByUrl("/qa")
          }

        }
      });
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
