package ParseEmail

import "testing"

func TestAddress(t *testing.T) {
	type args struct {
		emailAddr string
	}
	tests := []struct {
		name         string
		args         args
		wantUsername string
		wantDomain   string
		errValue     error
	}{
		{"john@toebes.com", args{"john@toebes.com"}, "john", "toebes.com", nil},
		{"john+tag1@toebes.com", args{"john+tag1@toebes.com"}, "john", "toebes.com", nil},
		{"john@@toebes.com", args{"john@@toebes.com"}, "", "", ErrBadEmailSyntax},
		{"john@@@toebes.com", args{"john@@@toebes.com"}, "", "", ErrBadEmailSyntax},
		{"@@", args{"@@"}, "", "", ErrBadEmailSyntax},
		{"@toebes.com", args{"@toebes.com"}, "", "toebes.com", ErrBadUsername},
		{"john@", args{"john@"}, "john", "", ErrBadDomain},
		{"~~john@toebes.com", args{"~~john@toebes.com"}, "~~john", "toebes.com", ErrBadUsername},
		{"john@t~oebes.com", args{"john@t~oebes.com"}, "john", "t~oebes.com", ErrBadDomain},
		{"<null string>", args{""}, "", "", ErrMissingAt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUsername, gotDomain, err := Address(tt.args.emailAddr)
			if err != tt.errValue {
				t.Errorf("ParseEmailAddress() error = %v, wantErr %v", err, tt.errValue)
				return
			}
			if gotUsername != tt.wantUsername {
				t.Errorf("ParseEmailAddress() gotUsername = %v, want %v", gotUsername, tt.wantUsername)
			}
			if gotDomain != tt.wantDomain {
				t.Errorf("ParseEmailAddress() gotDomain = %v, want %v", gotDomain, tt.wantDomain)
			}
		})
	}
}
