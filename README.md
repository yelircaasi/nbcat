# nbcat-go

CLI tool to display the contents of a Jupyter notebook (.ipynb), similar to 
`cat` and `glow`, but for Jupyter notebooks (and admittedly much less fancy).

## Usage

```sh
nbcat my_notebook.ipynb
```

## Roadmap

- [x] Basic working prototype
- [ ] Test suite
- [ ] refactor parameters into constants 
- [ ] move declarations into wide scope to avoid unnecessary redefinition 
      inside functions
- [ ] Code review from someone with more than 2 days of Golang experience
- [ ] change cell number format to look more authentic
