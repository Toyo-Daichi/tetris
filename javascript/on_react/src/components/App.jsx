import React from 'react'
//
import './App.css'

const App = () => {

  const Header = () => {
    return (
      <div className='header'>
        <p>Reactによるテトリス作成</p>
      </div>
    )
  }

  const BoardTable = () => {
    return (
      <>
        <table className='board-table'>
          <tbody>
            <Tetris width="12" height="12" />
          </tbody>
        </table>
        <div className='board-result'>
          result
        </div>
      </>
    )
  }

  class Tetris extends React.Component {
    /*
    width = 12
    height = 12
    */

    constructor(props){
      super(props)
      this.width = props.width
      this.height = props.height
    }

    render(){
      return (
        <div className='tetris-block'>
          {this.width}, {this.height}
        </div>
      )
    }
  }

  return (
    <div className='container'>
      <Header />
      <div className='board'>
        <BoardTable />
      </div>
    </div>
  )
}

export default App
