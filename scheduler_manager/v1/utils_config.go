package v1

import t "github.com/marvin-hansen/goSM/config_types"

func (c *ScheduleManager) addConfig(config *t.SchedulerConfig) {
	c.state.configMap[config.SchedulerID] = config
}

func (c *ScheduleManager) getConfig(configID string) *t.SchedulerConfig {
	exists := c.checkConfigExists(configID)
	if !exists {
		return nil
	}
	return c.state.configMap[configID]
}

func (c *ScheduleManager) deleteConfig(configID string) {
	exists := c.checkConfigExists(configID)
	if !exists { //doesn't exists, do nothing
		return
	}

	delete(c.state.configMap, configID)
}

func (c *ScheduleManager) deleteConfigMap() {
	c.state.configMap = make(map[string]*t.SchedulerConfig)
}

func (c *ScheduleManager) checkConfigExists(client string) (exists bool) {
	if _, ok := c.state.configMap[client]; ok {
		return true
	}
	return false
}

func (c *ScheduleManager) isConfigMapEmpty() bool {
	if len(c.state.configMap) == 0 {
		return true
	}
	return false
}
