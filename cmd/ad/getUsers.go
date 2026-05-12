package ad

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/go-ldap/ldap/v3"
	"github.com/jercle/cloudini/lib"
)

func GetAllADUsers(options lib.ADDomainConfig) (users []ADUser) {

	searchFilter := "(&(objectClass=user))"

	l, err := ldap.DialURL("ldap://" + options.DomainController + ":389")
	lib.CheckFatalError(err)
	defer l.Close()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	err = l.StartTLS(tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = l.Bind(options.BindUser, options.BindPwd)
	lib.CheckFatalError(err)

	searchRequest := ldap.NewSearchRequest(
		options.BaseSearchDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		searchFilter, // The filter to apply
		[]string{},   // A list attributes to retrieve
		nil,
	)

	sr, err := l.Search(searchRequest)
	lib.CheckFatalError(err)
	// lib.JsonMarshalAndPrint(sr)

	fmt.Println(sr)
	os.Exit(0)

	usrs := SearchResponseUserTransform(sr)

	for _, usr := range usrs {
		curr := usr
		curr.Domain = options.Domain
		users = append(users, curr)
	}

	return
}

func GetAllADUsersForAllConfiguredDomains(options lib.ActiveDirectoryConfig) (users []ADUser) {

	for _, dConfig := range options.Domains {
		dUsers := GetAllADUsers(dConfig)
		users = append(users, dUsers...)
	}
	return
}
