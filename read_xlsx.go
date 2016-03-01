package main

import (
  "fmt"
  // "bytes"
  "github.com/tealeg/xlsx"
  "github.com/wiggitywalt/xlsx_read/models"
  "github.com/wiggitywalt/xlsx_read/mongoService"
  "github.com/wiggitywalt/xlsx_read/rostrService"
  "strings"
  // "sync"
)

func main() {
  // var wg sync.WaitGroup
  excelFileName := "project_actuals_01.xlsx"
  xlFile, err := xlsx.OpenFile(excelFileName)

  if err != nil {
    fmt.Println("error!")
  }

  // var emps []string //list of employees, dupes to be removed
  // var wbs_projects []string //
  //sheet 0 has person info
  //Sheet 1 has project info

  for idx, sheet := range xlFile.Sheets {
    if idx == 1{
      for idx, row := range sheet.Rows {
        if idx > 0{

          fmt.Println("0", row.Cells[0], "1", row.Cells[1],"2", row.Cells[2],"3", row.Cells[3],"4", row.Cells[4],"5",row.Cells[5] )
          // if row.Cells[8].String() != "Employee"{
            // var thisuser = findAPerson(row.Cells[8].String())
            // var wbs = row.Cells[3].String()
            // var startdate = row.Cells[13].String()
            // insertPersonProject(thisuser, wbs , startdate)
            //userid string, wbs string, startdate time.Time
            // fmt.Println(findAPerson(row.Cells[8].String()), row.Cells[3].String(), row.Cells[13].String())
          // }
        }

        // for idx, cell := range row.Cells {
        //   fmt.Println(cell)
        //   if idx == 8 && len(cell.String()) > 0 {
        //     emps = append(emps, cell.String())
        //   }
        //   if idx == 5 && len(cell.String()) > 1{
        //     wbs_projects = append(wbs_projects, cell.String())
        //   }
        // }
      }
    }
  }


  // dedupedEmps := removeDuplicates(emps)
  // dedupedWBS := removeDuplicates(wbs_projects)

  // for _, emp := range dedupedEmps{
  //   if emp != "Employee"{
      // fmt.Println(emp)
    //   fmt.Println(findAPerson(emp))
    // }

  }
  // processEmps(dedupedEmps, &wg)
  // processWbss(dedupedWBS, &wg)
  // wg.Wait()
  // fmt.Println(reflect.TypeOf(dedupedEmps))
  // fmt.Println(reflect.TypeOf(dedupedWBS))
// }

// func processEmps(emps []string, wg *sync.WaitGroup){
//   wg.Add(1)
//   for _, emp := range emps{
//     if emp != "Employee"{
//       fmt.Println(emp)
//       insertPersonRecord(emp)
//         // emp1 := getRostrInfoByEmp(emp)
//         // //the length of the Rostr reply should be way larger than 75. less than that means there isn't a record.
//         // if len(emp1) > 75{
//         //  result := bytes.Index(bytes.ToLower(emp1), []byte("email"))
//         //  end_result := bytes.Index(bytes.ToLower(emp1), []byte("@disney.com"))
//         //  email_result := string(emp1[result:end_result+11])
//         //  email01 := strings.Replace(email_result, "email","", -1)
//         //  first3 := email01[0:3]
//         //  fmt.Println("User ", emp, " email address: ", strings.ToLower(stripchars(email01,first3)))
//         //  insertPersonRecord(emp,strings.ToLower(stripchars(email01,first3)))
//         // } else {
//         //  email01 := "No Email Found"
//         //  fmt.Println("User ", emp, " email address: ", email01)
//         //  insertPersonRecord(emp,email01)
//         // }
//     }
//   }
//   wg.Done()
// }
// func processWbss(wbss []string, wg *sync.WaitGroup){
//   wg.Add(1)
//   for _, wbs := range wbss{
//     var test_string = strings.Split(wbs, "~")
//     if len(test_string) > 1{
//       insertProjectRecord(test_string[0], test_string[1])
//       fmt.Println("WBS is ", test_string[0], " project is ", test_string[1])
//     }
//   }
//   wg.Done()
// }

func findAPerson(username string) string {
    session := mongoService.CreateMongoSession()
    defer session.Close()
    aperson := models.GetPerson(session, username)
    person_id := aperson.ID
    result_id := person_id.Hex()
    return result_id
}


func removeDuplicates(elements []string) []string {
  // Use map to record duplicates as we find them.
  encountered := map[string]bool{}
  result := []string{}

  for v := range elements {
    if encountered[elements[v]] == true {
      // Do not add duplicate.
    } else {
      // Record this element as an encountered element.
      encountered[elements[v]] = true
      // Append to result slice.
      result = append(result, elements[v])
    }
  }
  // Return the new slice.
  return result
}

func insertPersonProject(myid string, wbs string, startdate string){
  session := mongoService.CreateMongoSession()
  defer session.Close()

  fmt.Println(myid, wbs, startdate)
  models.AddPersonProject(session, models.Personproject{Userid: myid, WBS:wbs, Startdate: startdate})
}

func insertPersonRecord(person_name string){
  session := mongoService.CreateMongoSession()
  defer session.Close()
  models.AddPerson(session, models.Person{Name:person_name})
}

func insertProjectRecord(wbs string, project_name string){
  session := mongoService.CreateMongoSession()
  defer session.Close()
  models.AddProject(session, models.Project{WBS: wbs, Name:project_name})
}

func getRostrInfoByEmp(emp string) []byte {
  rostr_info := rostr.RetrieveInfo(emp)
  return rostr_info
}

func stripchars(str, chr string) string {
  return strings.Map(func(r rune) rune {
    if strings.IndexRune(chr, r) < 0 {
      return r
    }
    return -1
  }, str)
}
