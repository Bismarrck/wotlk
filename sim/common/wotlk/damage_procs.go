package wotlk

import (
	"time"

	"github.com/wowsims/wotlk/sim/core"
)

type ProcDamageEffect struct {
	ID      int32
	Trigger core.ProcTrigger

	School core.SpellSchool
	MinDmg float64
	MaxDmg float64
}

func newProcDamageEffect(config ProcDamageEffect) {
	core.NewItemEffect(config.ID, func(agent core.Agent) {
		character := agent.GetCharacter()

		minDmg := config.MinDmg
		maxDmg := config.MaxDmg
		damageSpell := character.RegisterSpell(core.SpellConfig{
			ActionID:    core.ActionID{ItemID: config.ID},
			SpellSchool: config.School,
			ProcMask:    core.ProcMaskEmpty,

			DamageMultiplier: 1,
			CritMultiplier:   character.DefaultSpellCritMultiplier(),
			ThreatMultiplier: 1,

			ApplyEffects: func(sim *core.Simulation, target *core.Unit, spell *core.Spell) {
				spell.CalcAndDealDamage(sim, target, sim.Roll(minDmg, maxDmg), spell.OutcomeMagicHitAndCrit)
			},
		})

		triggerConfig := config.Trigger
		triggerConfig.Handler = func(sim *core.Simulation, _ *core.Spell, _ *core.SpellResult) {
			damageSpell.Cast(sim, character.CurrentTarget)
		}
		core.MakeProcTriggerAura(&character.Unit, triggerConfig)
	})
}

func init() {
	core.AddEffectsToTest = false

	newProcDamageEffect(ProcDamageEffect{
		ID: 37064,
		Trigger: core.ProcTrigger{
			Name:       "Vestige of Haldor",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMeleeOrRanged,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolFire,
		MinDmg: 1024,
		MaxDmg: 1536,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 37264,
		Trigger: core.ProcTrigger{
			Name:       "Pendulum of Telluric Currents",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskSpellDamage,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 1168,
		MaxDmg: 1752,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 39889,
		Trigger: core.ProcTrigger{
			Name:       "Horn of Agent Fury",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMeleeOrRanged,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolHoly,
		MinDmg: 1024,
		MaxDmg: 1536,
	})

	core.AddEffectsToTest = true

	newProcDamageEffect(ProcDamageEffect{
		ID: 40371,
		Trigger: core.ProcTrigger{
			Name:       "Bandit's Insignia",
			Callback:   core.CallbackOnSpellHitDealt,
			ProcMask:   core.ProcMaskMeleeOrRanged,
			Outcome:    core.OutcomeLanded,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolArcane,
		MinDmg: 1504,
		MaxDmg: 2256,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 40373,
		Trigger: core.ProcTrigger{
			Name:       "Extract of Necromantic Power",
			Callback:   core.CallbackOnPeriodicDamageDealt,
			Harmful:    true,
			ProcChance: 0.10,
			ICD:        time.Second * 15,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 788,
		MaxDmg: 1312,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 42990,
		Trigger: core.ProcTrigger{
			Name:       "DMC Death",
			Callback:   core.CallbackOnSpellHitDealt | core.CallbackOnPeriodicDamageDealt,
			Harmful:    true,
			ProcChance: 0.15,
			ICD:        time.Second * 45,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 1750,
		MaxDmg: 2250,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 50415,
		Trigger: core.ProcTrigger{
			Name:     "Bryntroll, the Bone Arbiter (N) Proc",
			Callback: core.CallbackOnSpellHitDealt,
			ProcMask: core.ProcMaskMeleeOrRanged,
			Harmful:  true,
			PPM:      1.0,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 2138,
		MaxDmg: 2362,
	})

	newProcDamageEffect(ProcDamageEffect{
		ID: 50709,
		Trigger: core.ProcTrigger{
			Name:     "Bryntroll, the Bone Arbiter (H) Proc",
			Callback: core.CallbackOnSpellHitDealt,
			ProcMask: core.ProcMaskMeleeOrRanged,
			Harmful:  true,
			PPM:      1.0,
		},
		School: core.SpellSchoolShadow,
		MinDmg: 2412,
		MaxDmg: 2664,
	})
}
