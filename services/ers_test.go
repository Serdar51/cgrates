/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/
package services

/*
//TestEventReaderSCoverage for cover testing
func TestEventReaderSCoverage(t *testing.T) {
	cfg := config.NewDefaultCGRConfig()

	utils.Logger, _ = utils.Newlogger(utils.MetaSysLog, cfg.GeneralCfg().NodeID)
	utils.Logger.SetLogLevel(7)
	cfg.SessionSCfg().Enabled = true
	filterSChan := make(chan *engine.FilterS, 1)
	filterSChan <- nil
	shdChan := utils.NewSyncedChan()
	shdWg := new(sync.WaitGroup)
	server := cores.NewServer(nil)
	srvMngr := servmanager.NewServiceManager(cfg, shdChan, shdWg)
	srvDep := map[string]*sync.WaitGroup{utils.DataDB: new(sync.WaitGroup)}
	anz := NewAnalyzerService(cfg, server, filterSChan, shdChan, make(chan rpcclient.ClientConnector, 1), srvDep)
	db := NewDataDBService(cfg, nil, srvDep)
	sS := NewSessionService(cfg, db, server, make(chan rpcclient.ClientConnector, 1), shdChan, nil, nil, anz, srvDep)
	attrS := NewEventReaderService(cfg, filterSChan, shdChan, nil, srvDep)
	engine.NewConnManager(cfg, nil)
	srvMngr.AddServices(attrS, sS,
		NewLoaderService(cfg, db, filterSChan, server, make(chan rpcclient.ClientConnector, 1), nil, anz, srvDep), db)
	if err := srvMngr.StartServices(); err != nil {
		t.Error(err)
	}
	if attrS.IsRunning() {
		t.Errorf("Expected service to be down")
	}

		var reply string
		if err := cfg.V1ReloadConfig(&config.ReloadArgs{
			Path:    path.Join("/usr", "share", "cgrates", "conf", "samples", "ers_reload", "internal"),
			Section: config.ERsJson,
		}, &reply); err != nil {
			t.Error(err)
		} else if reply != utils.OK {
			t.Errorf("Expecting OK ,received %s", reply)
		}
		time.Sleep(10 * time.Millisecond) //need to switch to gorutine
		if !attrS.IsRunning() {
			t.Errorf("Expected service to be running")
		}
		err := attrS.Start()
		if err == nil || err != utils.ErrServiceAlreadyRunning {
			t.Errorf("\nExpecting <%+v>,\n Received <%+v>", utils.ErrServiceAlreadyRunning, err)
		}
		err = attrS.Reload()
		if err != nil {
			t.Errorf("\nExpecting <nil>,\n Received <%+v>", err)
		}
		cfg.ERsCfg().Enabled = false
		cfg.GetReloadChan(config.ERsJson) <- struct{}{}
		time.Sleep(10 * time.Millisecond)
		if attrS.IsRunning() {
			t.Errorf("Expected service to be down")
		}
		shdChan.CloseOnce()
		time.Sleep(10 * time.Millisecond)

}
*/