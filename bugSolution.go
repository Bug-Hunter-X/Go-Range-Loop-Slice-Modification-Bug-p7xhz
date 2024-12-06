func myFunc(a []int) {
  for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
    // Modification here will correctly change the slice
    a[i] = a[i] * 2
  }
} 