import { Routes, Route } from 'react-router-dom';

import Layout from './components/Layout';

import IssuesPage from './pages/IssuesPage';
import EditIssuePage from './pages/EditIssuePage';
import CreateIssuePage from './pages/CreateIssuePage';

import EditArticlePage from './pages/EditArticlePage';
import CreateArticlePage from './pages/CreateArticlePage';

import CouncilsPage from './pages/CouncilsPage';
import EditCouncilPage from './pages/EditCouncil';
import CreateCouncilPage from './pages/CreateCouncilPage';
import EditReductorPage from './pages/EditReductorPage';
import CreateReductorPage from './pages/CreateReductorPage';



import FilesPage from './pages/FilesPage';

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";


const queryClient = new QueryClient();


function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Routes>
        <Route path="/" element={<Layout/>}>
          <Route index element={<IssuesPage/>} />

          <Route path="issues" element={<IssuesPage/>} />
          <Route path="issues/:id" element={<EditIssuePage/>} />
          <Route path="issues/create" element={<CreateIssuePage/>} />

          <Route path="issues/article/:id" element={<EditArticlePage/>} />
          <Route path="issues/article/create/:id" element={<CreateArticlePage/>} />

          <Route path="reductors" element={<CouncilsPage/>} />
          <Route path="councils/:id" element={<EditCouncilPage/>} />
          <Route path="councils/create" element={<CreateCouncilPage/>} />
          <Route path="reductors/:id" element={<EditReductorPage/>} />
          <Route path="reductors/create" element={<CreateReductorPage/>} />

          <Route path="files" element={<FilesPage/>} />
        </Route>
      </Routes>
    </QueryClientProvider>
  )
}

export default App

